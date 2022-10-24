package br.unifei.imc.entities.users;

import lombok.Getter;
import lombok.Setter;

public interface User {
    @Getter
    @Setter
    String Name = null;
    @Getter
    @Setter
    String Email = null;
}
