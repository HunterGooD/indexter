<p align="center">
  <a href="https://github.com/HunterGooD/indexter" rel="noopener">
 <img width=415px height=115x src="docs/images/logo.png" alt="Project logo"></a>
</p>

<h3 align="center">Indexter indexed you directory</h3>

---

<p align="center"> 
  Индексация файлов txt, md в директории и поиск по словам
    <br> 
</p>

## 📝 Содержание

- [О проекте](#about)
- [Быстрый старт](#quick_start)
- [Установка](#install)

## 🧐 О проекте <a name = "about"></a>

Создает индексный файл, обеспечивающий быстрый поиск файлов по
словам. Функции:

- добавление новой индексируемой папки;
- поиск файлов в индексной директории по словам.

### 🚀 Быстрый старт <a name = "quick_start"></a>

Скачать бинарный файл для разных ОС можно на странице [releases](https://github.com/HunterGooD/indexter/releases).

### Установка <a name="install"></a>

Для начала скопировать репозиторий в вашу директорию для приложения.

Для этого нужно выполнить следующую команду

```bash
  $ git clone https://github.com/HunterGooD/indexter.git
```

После клонирования репозитория нужно скомпилиривать проект для этого нужно иметь компилятор `go` версии `1.17+` и `make`

```bash
 $ make build
```

После выполнения команды появится директория `dist` в которой содержиться бинарный файл программы. 
