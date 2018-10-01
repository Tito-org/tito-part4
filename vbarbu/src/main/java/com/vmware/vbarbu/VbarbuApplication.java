package com.vmware.vbarbu;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.web.servlet.ServletComponentScan;

@ServletComponentScan
@SpringBootApplication
public class VbarbuApplication {

    public static void main(String[] args) {
        SpringApplication.run(VbarbuApplication.class, args);
    }
}
