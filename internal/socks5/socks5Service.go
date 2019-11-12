package socks5



type Sock5Server struct {
	quit chan struct{}
}
func (s *Sock5Server) Start(url string){
	go func() {
		for {
			select {
			case <-s.quit:
				.Conn.Close()
				break
			default:
				go func(u string, s *Server) {
					if err := server.ListenAndServe("tcp", u); err != nil {
						panic(err)
					}
				}(url, server)

			}
		}
	}()
}
func (s *Sock5Server) Stop() {
	close(s.quit)
}
