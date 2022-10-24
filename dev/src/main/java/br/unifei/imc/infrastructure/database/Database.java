package br.unifei.imc.infrastructure.database;

public interface Database {
    public void connect();

    public void disconnect();

    public void insert(String table, String[] columns, String[] values);

    public void update(String table, String[] columns, String[] values, String where);

    public void delete(String table, String where);

    public void select(String table, String[] columns, String where);

    public void createTable(String table, String[] columns, String[] types);

    public void createDatabase(String database);

    public void useDatabase(String database);
}
