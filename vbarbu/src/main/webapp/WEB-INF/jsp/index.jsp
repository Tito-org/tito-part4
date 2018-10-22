<%@ page import="java.awt.*" %>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<%@ page language="java" contentType="text/html; charset=ISO-8859-1" pageEncoding="ISO-8859-1"%>

<html>

<style>
    .button-ret {
        background-color: transparent;
        color: black;
        transition-duration: 0.4s;
        padding: 15px 32px;
        border: 2px solid black;
        top: 90%;
        cursor: pointer;
        border-radius: 8px;
    }

    .button-ret:hover {
        background-color: black;
        color: white;
    }

    table {
        border-radius: 3px;
    }
</style>

<body style="background-image: linear-gradient(#BEEEFF, #3FCDFF);">

<div align="center">
    <table border="3">
        <th>Model</th>
        <th>Marque</th>
        <th>GazLevel (in %)</th>
        <th>Location (in Km)</th>
        <th>Image</th>
        <th>Disponibility</th>
        <c:forEach items="${pictureList}" var="picture">
            <tr>
                <td>${picture.model}</td>
                <td>${picture.trade}</td>
                <td>${picture.gazLevel}</td>
                <td>${picture.location}</td>

                <td><img src="data:image/jpg;base64,${picture.base64Image}" width="240" height="200"/></td>
                <c:if test="${picture.book == false}">
                        <td> Available </td>
                </c:if>
                <c:if test="${picture.book == true}">
                    <td> Not Available </td>
                </c:if>
            </tr>
        </c:forEach>
    </table>
</div>
<form action="https://shwrfr.com/tito/">
    <input type="submit" value="Return" class="button-ret">
</form>
</body>
</html>