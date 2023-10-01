package logging

type Logger interface {
    Info(m string)
    Infow(m string, args ...interface{})
    Warning(m string)
    Warningw(m string, args ...interface{})
    Error(m string)
    Errorw(m string, args ...interface{})
    Fatal(m string)
    Fatalw(m string, args ...interface{})
}
