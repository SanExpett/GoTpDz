Описание задачи
Тут раелизована утилита, с помощью которой можно вывести или отфильтровать повторяющиеся строки в файле (аналог UNIX утилиты uniq). Причём повторяющиеся входные строки не должны распозноваться, если они не следуют строго друг за другом. Сама утилита имеет набор параметров, которые она поддерживает.

Параметры
-с - подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.

-d - вывести только те строки, которые повторились во входных данных.

-u - вывести только те строки, которые не повторились во входных данных.

-f num_fields - не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.

-s num_chars - не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля).

-i - не учитывать регистр букв.

Все параметры опциональны. Поведения утилиты без параметров -- простой вывод уникальных строк из входных данных.

Параметры c, d, u взаимозаменяемы. Необходимо учитывать, что параллельно эти параметры не имеют никакого смысла. При передаче одного вместе с другим нужно отобразить пользователю правильное использование утилиты

Если не передан input_file, то входным потоком считать stdin

Если не передан output_file, то выходным потоком считать stdout


Пример использования:

Запуск программы:
go run main.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]

Запуск тестов:
go test ./...