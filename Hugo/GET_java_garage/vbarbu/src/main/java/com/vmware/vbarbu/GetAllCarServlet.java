package com.vmware.vbarbu;

import java.io.IOException;
import java.util.List;
import javax.servlet.RequestDispatcher;
import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import static com.vmware.vbarbu.CarDAO.displayPicture;

@WebServlet("/getcar")
public class GetAllCarServlet extends HttpServlet {
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        List<Car> list = null;
        list = CarDAO.displayPicture();

        request.setAttribute("pictureList", list);

        RequestDispatcher dispatcher = request.getServletContext().getRequestDispatcher("/WEB-INF/jsp/index.jsp");
        dispatcher.forward(request, response);

        //String page = "/index.jsp";
        //RequestDispatcher requestDispatcher = request.getRequestDispatcher(page);
        //requestDispatcher.forward(request, response);

    }
}
