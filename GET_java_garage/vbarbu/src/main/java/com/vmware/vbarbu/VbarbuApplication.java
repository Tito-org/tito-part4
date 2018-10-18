package com.vmware.vbarbu;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.web.servlet.ServletComponentScan;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@ServletComponentScan
@SpringBootApplication
@EnableSwagger2
public class VbarbuApplication {

    public static void main(String[] args) {
        SpringApplication.run(VbarbuApplication.class, args);
    }
}
