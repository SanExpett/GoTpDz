package main

import (
	"log"
	"sort"
	"strconv"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	channels := make([]chan interface{}, len(cmds)+1)
	for i := range channels {
		channels[i] = make(chan interface{})
	}

	wg := sync.WaitGroup{}
	wg.Add(len(cmds))

	for i := range cmds {
		go func(i int) {
			defer wg.Done()

			cmds[i](channels[i], channels[i+1])
			close(channels[i+1])
		}(i)
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	// 	in - string
	// 	out - User
	processedEmails := make(map[uint64]bool)
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	for email := range in {
		emailStrConverted, ok := email.(string)
		if !ok {
			continue
		}

		wg.Add(1)
		go func(emailStr string) {
			defer wg.Done()

			user := GetUser(emailStr)

			mu.Lock()
			if !processedEmails[user.ID] {
				processedEmails[user.ID] = true

				out <- user
			}
			mu.Unlock()
		}(emailStrConverted)
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	// 	in - User
	// 	out - MsgID
	wg := sync.WaitGroup{}
	batchSize := 2
	batch := make([]User, 0, batchSize)

	processBatch := func(batch []User) {
		defer wg.Done()

		messagesBatch, err := GetMessages(batch...)
		if err != nil {
			log.Printf("Error getting messages for batch of users: %s\n", err)

			return
		}

		for _, message := range messagesBatch {
			out <- message
		}
	}

	for user := range in {
		if userConverted, ok := user.(User); ok {
			batch = append(batch, userConverted)
		}

		if len(batch) == batchSize {
			wg.Add(1)
			go func(batch []User) {
				processBatch(batch)
			}(batch)

			batch = make([]User, 0, batchSize)
		}
	}

	if len(batch) > 0 {
		wg.Add(1)
		go func(batch []User) {
			processBatch(batch)
		}(batch)
	}

	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData
	var wg sync.WaitGroup

	maxRequestCount := 5
	concurrency := make(chan struct{}, maxRequestCount)

	for message := range in {
		msgIDConverted, ok := message.(MsgID)
		if !ok {
			continue
		}

		//msgID := message.(MsgID)

		concurrency <- struct{}{}

		wg.Add(1)
		go func(msgID MsgID) {
			defer func() {
				wg.Done()
				<-concurrency
			}()

			isSpam, err := HasSpam(msgID)
			if err != nil {
				log.Printf("Error checking spam for message %d: %s\n", msgID, err)

				return
			}

			msgData := MsgData{ID: msgID, HasSpam: isSpam}

			out <- msgData
		}(msgIDConverted)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string
	var msgDataSlice []MsgData

	for msgData := range in {
		if msgDataConverted, ok := msgData.(MsgData); ok {
			msgDataSlice = append(msgDataSlice, msgDataConverted)
		}
	}

	sort.Slice(msgDataSlice, func(i, j int) bool {
		if msgDataSlice[i].HasSpam && !msgDataSlice[j].HasSpam {
			return true
		}

		if msgDataSlice[i].HasSpam == msgDataSlice[j].HasSpam {
			return msgDataSlice[i].ID < msgDataSlice[j].ID
		}

		return false
	})

	for _, msgData := range msgDataSlice {
		msgDataStr := strconv.FormatBool(msgData.HasSpam) + " " + strconv.FormatUint(uint64(msgData.ID), 10)
		out <- msgDataStr
	}
}

func main() {
	RunPipeline(
		SelectUsers,
		SelectMessages,
		CheckSpam,
		CombineResults,
	)
}
