# goback
Based on WM_COPYDATA, realize two-way communication between go language and c++, it is very convenient to use ralang to develop background service under windows platform, c++ is responsible for UI interface

# platform: windows

# example

```
// MainWnd is the name of the UI main window
obj := goback.Regist("MainWnd")
go func() {
	for {
		_, ok := <-obj.BufCh
		if !ok {
			break
		}
	}
        close(obj.BufCh)
}()
goback.Wait()
// end
```

# notes

WM_COPYDATA communication is synchronously blocked, and can only send data through SendMessage, so the project can only meet the demand in the case of small business volume.
