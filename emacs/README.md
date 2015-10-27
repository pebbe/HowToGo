## Go & Emacs

Some useful hints for editing Go source code in Emacs.

Tested with GNU Emacs 23.3.2

----

Build errors for Go files that import package "C" won't parse
correct in Emacs' compilation buffer. To fix this, follow these
steps:

1. Type: *M-x* `customize-variable` *RET* `compilation-error-regexp-alist` *RET*
2. At the bottom of the list, click: INS 
3. Select: ValueMenu: Predefined symbol: nil
4. Select: ValueMenu: Error specification: `("^\\([^ \t\n:][^:\n]*go\\):\\([0-9][0-9]*\\)\\[[^]]*go:[0-9][0-9]*\\]:" 1 2 nil 2)`
5. Select: State: Save for future sessions

----

Some useful functions to add to your `~/.emacs` file.

```lisp
(add-hook
 'go-mode-hook
 '(lambda ()

    <font color="#f00">;; Imenu & Speedbar</font>

    (setq imenu-generic-expression
          '(("type" "^type *\\([^ \t\n\r\f]*\\)" 1)
            ("func" "^func *\\(.*\\) {" 1)))
    (imenu-add-to-menubar "Index")

    <font color="#f00">;; Outline mode</font>

    <font color="#f00">;; Level 3: <span class="code">//.</span>  use this to devide the file into major sections</font>
    <font color="#f00">;; Level 4: <span class="code">//</span>   followed by at least two characters</font>
    <font color="#f00">;; Level 4: <span class="code">pack</span>age</font>
    <font color="#f00">;; Level 4: <span class="code">impo</span>rt</font>
    <font color="#f00">;; Level 4: <span class="code">cons</span>t</font>
    <font color="#f00">;; Level 4: <span class="code">var</span>  followed by at least one character</font>
    <font color="#f00">;; Level 4: <span class="code">type</span></font>
    <font color="#f00">;; Level 4: <span class="code">func</span></font>
    <font color="#f00">;; Level 5 and above: tab-indented lines with at least five characters</font>
    (make-local-variable 'outline-regexp)
    (setq outline-regexp "//\\.\\|//[^\r\n\f][^\r\n\f]\\|pack\\|func\\|impo\\|cons\\|var[^\r\n\f]\\|type\\|\t\t*[^\r\n\f]\\{4\\}")
    (outline-minor-mode 1)
    (local-set-key "\M-a" 'outline-previous-visible-heading)
    (local-set-key "\M-e" 'outline-next-visible-heading)

    <font color="#f00">;; Menu bar</font>

    (require 'easymenu)
    (defconst go-hooked-menu
      '("Go tools"
        ["Go run buffer" go t]
        ["Go build buffer" go-build t]
        ["Go build directory" go-build-dir t]
        ["Go reformat buffer" go-fmt-buffer t]
        "---"
        ["Go check buffer" go-fix-buffer t]
        "---"
        ["Go install package" go-install-package t]
        ["Go test package" go-test-package t]))
    (easy-menu-define
      go-added-menu
      (current-local-map)
      "Go tools"
      go-hooked-menu)

    <font color="#f00">;; Other</font>

    (setq tab-width 4)
    (setq show-trailing-whitespace t)

    ))

<font color="#f00">;; helper variable</font>
(defvar hook-go-pkg nil
  "History variable for `go-install-package' and `go-test-package'.")

<font color="#f00">;; helper function</font>
(defun go ()
  "run current buffer"
  (interactive)
  (compile (concat "go run \"" (buffer-file-name) "\"")))

<font color="#f00">;; helper function</font>
(defun go-build ()
  "build current buffer"
  (interactive)
  (compile (concat "go build  \"" (buffer-file-name) "\"")))

<font color="#f00">;; helper function</font>
(defun go-build-dir ()
  "build current directory"
  (interactive)
  (compile "go build ."))

<font color="#f00">;; helper function</font>
(defun go-fmt-buffer ()
  "run gofmt on current buffer"
  (interactive)
  (if buffer-read-only
    (progn
      (ding)
      (message "Buffer is read only"))
    (let ((p (line-number-at-pos))
          (filename (buffer-file-name))
          (old-max-mini-window-height max-mini-window-height))
      (show-all)
      (if (get-buffer "*Go Reformat Errors*")
          (progn
            (delete-windows-on "*Go Reformat Errors*")
            (kill-buffer "*Go Reformat Errors*")))
      (setq max-mini-window-height 1)
      (if (= 0 (shell-command-on-region (point-min) (point-max) "gofmt" "*Go Reformat Output*" nil "*Go Reformat Errors*" t))
          (progn
            (erase-buffer)
            (insert-buffer-substring "*Go Reformat Output*")
            (goto-char (point-min))
            (forward-line (1- p)))
        (with-current-buffer "*Go Reformat Errors*"
          (progn
            (goto-char (point-min))
            (while (re-search-forward "<standard input>" nil t)
              (replace-match filename))
            (goto-char (point-min))
            (compilation-mode))))
      (setq max-mini-window-height old-max-mini-window-height)
      (delete-windows-on "*Go Reformat Output*")
      (kill-buffer "*Go Reformat Output*"))))

<font color="#f00">;; helper function</font>
(defun go-fix-buffer ()
  "run gofix on current buffer"
  (interactive)
  (show-all)
  (shell-command-on-region (point-min) (point-max) "go tool fix -diff"))

<font color="#f00">;; helper function</font>
(defun go-install-package ()
  "install package"
  (interactive)
  (let
      ((pkg (read-from-minibuffer "Install package: " nil nil nil 'hook-go-pkg)))
    (if (not (string= pkg ""))
        (compile (concat "go install \"" pkg "\"")))))

<font color="#f00">;; helper function</font>
(defun go-test-package ()
  "test package"
  (interactive)
  (let
      ((pkg (read-from-minibuffer "Test package: " nil nil nil 'hook-go-pkg)))
    (if (not (string= pkg ""))
        (compile (concat "go test \"" pkg "\"")))))
```

----

To enable Go in the Speedbar:

1. Go to: Options → Customize Emacs → Specific Option...
2. Choose: speedbar-supported-extension-expressions
3. Add extension: .go
4. Click on: State → Save for Future Sessions

To start the Speedbar: *Alt-X* `speedbar` *Enter*

----

The following makes outline mode easy to use generally, not just when
editing Go. You may have to install outline-magic separately.

```lisp
<font color="#f00">;; Show/hide parts by repeated pressing f10</font>
(add-hook 'outline-minor-mode-hook
           (lambda ()
             (require 'outline-magic)
             (define-key outline-minor-mode-map [(f10)] 'outline-cycle)))
```
