package br.unifei.imc.infrastructure.database.SQLite;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

import br.unifei.imc.infrastructure.database.Database;

public class Sqlite implements Database {
    private Connection connection;

    public Sqlite() {
        this.connection = null;
    }

    @Override
    public void connect() {
        connection = null;

        try {
            Class.forName("org.sqlite.JDBC");
            connection = DriverManager.getConnection("jdbc:sqlite:imc.db");
        } catch (Exception e) {
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            System.exit(0);
        }

        if (connection != null) {
            System.out.println("Opened database successfully");
        } else {
            System.out.println("Failed to open database");
        }
    }

    @Override
    public void disconnect() {
        connection = null;

        System.out.println("Disconnected database successfully");
    }

    @Override
    public void insert(String table, String[] columns, String[] values) {
        String sql = "INSERT INTO " + table + " (";

        for (int i = 0; i < columns.length; i++) {
            sql += columns[i];

            if (i < columns.length - 1) {
                sql += ", ";
            }
        }

        sql += ") VALUES (";

        for (int i = 0; i < values.length; i++) {
            sql += values[i];

            if (i < values.length - 1) {
                sql += ", ";
            }
        }

        sql += ");";

        try {
            Statement statement = connection.createStatement();
            statement.executeUpdate(sql);
            statement.close();
        } catch (SQLException e) {
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            System.exit(0);
        }
    }

    @Override
    public void update(String table, String[] columns, String[] values, String where) {
        String sql = "UPDATE " + table + " SET ";

        for (int i = 0; i < columns.length; i++) {
            sql += columns[i] + " = " + values[i];

            if (i < columns.length - 1) {
                sql += ", ";
            }
        }

        sql += " WHERE " + where + ";";

        try {
            Statement statement = connection.createStatement();
            statement.executeUpdate(sql);
            statement.close();
        } catch (SQLException e) {
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            System.exit(0);
        }
    }

    @Override
    public void delete(String table, String where) {
        String sql = "DELETE FROM " + table + " WHERE " + where + ";";

        try {
            Statement statement = connection.createStatement();
            statement.executeUpdate(sql);
            statement.close();
        } catch (SQLException e) {
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            System.exit(0);
        }
    }

    @Override
    public void select(String table, String[] columns, String where) {
        String sql = "SELECT ";

        for (int i = 0; i < columns.length; i++) {
            sql += columns[i];

            if (i < columns.length - 1) {
                sql += ", ";
            }
        }

        sql += " FROM " + table + " WHERE " + where + ";";

        try {
            Statement statement = connection.createStatement();
            ResultSet resultSet = statement.executeQuery(sql);

            while (resultSet.next()) {
                for (int i = 0; i < columns.length; i++) {
                    System.out.println(columns[i] + ": " + resultSet.getString(columns[i]));
                }
            }

            resultSet.close();
            statement.close();
        } catch (SQLException e) {
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            System.exit(0);
        }
    }

    @Override
    public void createTable(String table, String[] columns, String[] types) {
        String sql = "CREATE TABLE " + table + " (";

        for (int i = 0; i < columns.length; i++) {
            sql += columns[i] + " " + types[i];

            if (i < columns.length - 1) {
                sql += ", ";
            }
        }

        sql += ");";

        try {
            Statement statement = connection.createStatement();
            statement.executeUpdate(sql);
            statement.close();
        } catch (SQLException e) {
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            System.exit(0);
        }
    }

    @Override
    public void createDatabase(String database) {
        String sql = "CREATE DATABASE " + database + ";";

        try {
            Statement statement = connection.createStatement();
            statement.executeUpdate(sql);
            statement.close();
        } catch (SQLException e) {
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            System.exit(0);
        }
    }

    @Override
    public void useDatabase(String database) {
        String sql = "USE " + database + ";";

        try {
            Statement statement = connection.createStatement();
            statement.executeUpdate(sql);
            statement.close();
        } catch (SQLException e) {
            System.err.println(e.getClass().getName() + ": " + e.getMessage());
            System.exit(0);
        }
    }

}
