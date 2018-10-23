<%@ page import="java.awt.*" %>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<%@ page language="java" contentType="text/html; charset=ISO-8859-1" pageEncoding="ISO-8859-1"%>

<html>

<style>
    .button-ret {
        background-color: rgba(158,158,158,0.5);
        color: black;
        transition-duration: 0.4s;
        padding: 15px 32px;
        border: 2px solid black;
        cursor: pointer;
        border-radius: 8px;
        position: absolute;
        top: 90%;
    }

    .button-ret:hover {
        background-color: black;
        color: white;
    }

    table {
        border-collapse: collapse;
    }

    th {
        border: 2px solid black;
        padding: 3px;
        background-color: rgba(74,74,74,0.5);
        text-align: center;
    }

    tr:nth-child(even) {
        background-color:  rgba(158,158,158,0.5);
    }

    tr:nth-child(odd):not(th) {
        background-color: rgba(191,191,191,0.5);

    }

    table, td {
        border: 1px solid black;
        text-align: center;
    }

    @media screen and (min-width: 650px){
        table {
            width: 70%;
        }

    }

    body {
        background-image: url("https://www.wallpaperup.com/uploads/wallpapers/2013/08/19/135949/00297a2528269ffcd01955cb6acb0a12.jpg");
        background-repeat: no-repeat;
        background-position: center;
        color: black;
        background-color: white;
        background-attachment: fixed;
        -webkit-background-size: cover;
        -moz-background-size: cover;
        -o-background-size: cover;
        background-size: cover;
    }

    .button-available {
        max-width: 12%;
        height: auto;
    }

    .button-not-available {
        max-width: 15%;
        height: auto;
    }

    .img-car {
        max-width: 60%;
        height: auto;
    }

</style>

<body>

<div>
    <table align="center">
        <th>Model</th>
        <th>Marque</th>
        <th>GazLevel (in %)</th>
        <th>Location (in Km)</th>
        <th>Image</th>
        <th>Disponibility</th>
        <c:forEach items="${pictureList}" var="picture">
            <tr>
                <td style="color: white;">${picture.model}</td>
                <td style="color: white;">${picture.trade}</td>
                <td style="color: white;">${picture.gazLevel}</td>
                <td style="color: white;">${picture.location}</td>

                <td><img src="data:image/jpg;base64,${picture.base64Image}" class="img-car"/></td>
                <c:if test="${picture.book == false}">
                    <td><img src="https://upload.wikimedia.org/wikipedia/commons/5/58/Disponible.png" class="button-available"></td>
                    <!-- <td> Available </td> -->
                </c:if>
                <c:if test="${picture.book == true}">
                    <td> <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/16/Deletion_icon.svg/600px-Deletion_icon.svg.png" class="button-not-available"> </td>
                   <!-- <td> Not Available </td> -->
                </c:if>
            </tr>
        </c:forEach>
    </table>
</div>
<!-- <form action="http://172.18.12.219:1234"> -->

<c:forEach items="${pictureList}" var="picture">
    <form action="http://${picture.ipenv}" method="GET">
</c:forEach>
    <input class="button-ret" type="submit" value="Return">
    </form>
</body>
</html>
