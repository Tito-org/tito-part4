package com.vmware.vbarbu;

public class Car {
    public String model;
    public String trade;
    public String base64Image;
    public String gazLevel;
    public String location;
    public Boolean book;
    public String ipenv;

    public Car() {
    }

    public String getModel() {
        return model;
    }

    public void setModel(String model) {
        this.model = model;
    }

    public String getTrade() {
        return trade;
    }

    public void setTrade(String trade) {
        this.trade = trade;
    }

    public String getBase64Image() {
        return base64Image;
    }

    public void setBase64Image(String base64Image) {
        this.base64Image = base64Image;
    }

    public String getGazLevel() {
        return gazLevel;
    }

    public void setGazLevel(String gazLevel) {
        this.gazLevel = gazLevel;
    }

    public String getLocation() {
        return location;
    }

    public void setLocation(String location) {
        this.location = location;
    }

    public Boolean getBook() {
        return book;
    }

    public void setBook(Boolean book) {
        this.book = book;
    }

    public String getIpenv() {
        return ipenv;
    }

    public void setIpenv(String ipenv) {
        this.ipenv = ipenv;
    }
}