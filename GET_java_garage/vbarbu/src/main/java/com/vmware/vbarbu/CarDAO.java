package com.vmware.vbarbu;

import java.sql.*;
import java.util.ArrayList;
import java.util.Base64;
import java.util.List;

public class CarDAO {
    public static List<Car> displayPicture() {

        List<Car> list = new ArrayList<Car>();

        try {

            Class.forName("com.mysql.jdbc.Driver");

            Connection conn = DriverManager.getConnection(""+ System.getenv("db_connector") +":"+ System.getenv("db_type") +"://"+ System.getenv("db_ip") +"/"+ System.getenv("db_name") +"?autoReconnect=true&useSSL=false", ""+ System.getenv("db_username") +"", ""+ System.getenv("db_password") +"");
            //Connection conn = DriverManager.getConnection("jdbc:mysql://172.18.12.219/Test?autoReconnect=true&useSSL=false", "root", "PASSWORD");

            String sql = "SELECT * FROM garage";
            PreparedStatement ps = conn.prepareStatement(sql);
            String ip = System.getenv("tito_ip");

            System.out.println("Ip = " + ip);

            ResultSet rs = ps.executeQuery();

            while (rs.next()) {
                //byte[] imageData = rs.getBytes("photo");
                //String base64Image = Base64.getEncoder().encodeToString(imageData);
                String model = rs.getString("model");
                String trade = rs.getString("marque");
                String base64Image = rs.getString("photo");
                String gazLevel = rs.getString("gazLevel");
                String location = rs.getString("location");
                Boolean book = rs.getBoolean("book");

                Car picture = new Car();
                picture.setBase64Image(base64Image);
                picture.setModel(model);
                picture.setTrade(trade);
                picture.setGazLevel(gazLevel);
                picture.setLocation(location);
                picture.setBook(book);
                picture.setIpenv(ip);
                list.add(picture);
            }
            conn.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return list;
    }
}