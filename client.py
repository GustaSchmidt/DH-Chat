#!/usr/bin/env python3

import socket
import os
import select
import sys

from util import cria_mensagem
from screen import Screen

USERNAME = ""

SERVER_IP = 'localhost'
# SERVER_PORT = 8888
SERVER_PORT = 8080



if __name__ == '__main__':
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as fd:
        fd.connect((SERVER_IP, SERVER_PORT))
        
        with Screen() as scr:
            msg = fd.recv(256).decode()
            scr.set_text(msg)
            uname = scr.get_text()
            scr._n_lines -= 1
            scr.set_text(uname, col=len(msg))
            fd.send(uname[:-1].encode('utf8'))

            confirm = fd.recv(3).decode()

            if confirm != "OK\x01":
                fd.send("exit\x01".encode('utf8'))
                scr.clear()
                exit(1)
            
            users_online = fd.recv(2).decode()
            scr.set_text(f'Users connected: {users_online}')

            while True:
                socket_list = [sys.stdin, fd]

                sockets, _, _ = select.select(socket_list, [], [])

                try:
                    for s in sockets:
                        if s == fd:
                            msg = s.recv(2048)
                            scr.set_text(msg.decode())
                            continue
                        else:
                            msg = scr.get_text()

                            fd.send(msg.encode())

                            msg = cria_mensagem("Você", msg, net=False)
                            scr.set_text(msg)
                except KeyboardInterrupt:
                    fd.send("exit\x01".encode('utf8'))
                    break
