## How to kill a process that was started with exec.Command().Start()

This kills just the top process, but leaves any subprocesses running:

```go
cmd := exec.Command( some_command )

cmd.Start()

cmd.Process.Kill()

cmd.Wait()
```

On **Linux**, this kills all the subprocesses too:

```go
cmd := exec.Command( some_command )

cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
cmd.Start()

pgid, _ := syscall.Getpgid(cmd.Process.Pid)
syscall.Kill(-pgid, 15)

cmd.Wait()
```
