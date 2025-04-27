Output demo.gif
Set Padding 10
#Set PlaybackSpeed 0.5 # Make output 2 times slower
Set WindowBar Colorful
Set TypingSpeed 100ms
Sleep 3s
Type "go run ./cmd/example/main.go foo bar biz bass buzz booboo baba"
Enter
Wait+Screen@15s  /booboo/
Sleep 1s
Set TypingSpeed 500ms
Space
Down
Down
Space
Down
Space
Up
Up
Space
Up
Space
Up
Space
Sleep 3s
Enter
Sleep 6s
Ctrl+C