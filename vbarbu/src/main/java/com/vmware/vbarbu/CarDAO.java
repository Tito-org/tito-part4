package com.vmware.vbarbu;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.util.ArrayList;
import java.util.Base64;
import java.util.List;

public class CarDAO {
    public static List<Car> displayPicture() {

        List<Car> list = new ArrayList<Car>();

        try {

            Class.forName("com.mysql.jdbc.Driver");

            Connection conn = DriverManager.getConnection("jdbc:mysql://172.18.12.219/Test?autoReconnect=true&useSSL=false", "root", "PASSWORD");

            String sql = "SELECT * FROM garage";
            PreparedStatement ps = conn.prepareStatement(sql);

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
                System.out.println(book);
                list.add(picture);
            }
            conn.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return list;
    }
}