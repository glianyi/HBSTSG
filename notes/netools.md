检测remote server的端口是否连通
nc -v host port
eg. nc -v 192.168.90.215 12345

linux network programming(https://medium.com/@chongye225/networking-with-c-cf15426cc270)
#### socket
To create a endpoint for networking communication. A new socket by itself is not particularly useful; though we’ve specified either a packet or stream-based connections it is not bound to a particular network interface or port. Instead socket returns a network descriptor that can be used with later calls to bind, listen and accept.
#### bind
The bind call associates an abstract socket with an actual network interface and port. It is possible to call bind on a TCP client however it's unusually unnecessary to specify the outgoing port.
#### listen
The listen call specifies the queue size for the number of incoming, unhandled connections i.e. that have not yet been assigned a network descriptor by accept Typical values for a high performance server are 128 or more.

#### why are server sockste passive
Server sockets do not actively try to connect to another host; instead they wait for incoming connections. Additionally, server sockets are not closed when the peer disconnects. Instead the client communicates with a separate active socket on the server that is specific to that connection.
Unique TCP connections are identified by the tuple (source ip, source port, destination ip, destination port) It is possible to have multiple connections from a web browser to the same server port (e.g. port 80) because the the source port on each arriving packet is unique. i.e. For a particular server port (e.g. port 80) there can be one passive server socket but multiple active sockets (one for each currently open connection) and the server's operating system maintains a lookup table that associates a unique tuple with active sockets, so that incoming packets can be correctly routed to the correct socket.

#### accept