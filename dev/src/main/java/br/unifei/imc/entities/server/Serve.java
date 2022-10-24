package br.unifei.imc.entities.server;

import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
public class Serve {
    String Name;
    String IP;
    int Port;
    int Status;

    public Serve(String name, String ip, int port, int status) {
        Name = name;
        IP = ip;
        Port = port;
        Status = status;
    }

}
