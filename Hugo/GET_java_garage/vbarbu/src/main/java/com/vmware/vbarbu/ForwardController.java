package com.vmware.vbarbu;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("/test")
public class ForwardController {
    @GetMapping
    String getView(Model model) {
        model.addAttribute("msg", "Hello there, This message has been injected from the controller method");
        return "forward:/getcar";
    }
}