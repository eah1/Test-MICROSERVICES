package com.challenge.microservice.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@RestController
public class HealthCheckGetController {

    @GetMapping(value = "/health")
    public HashMap<String, String> handlerHealth() {
        HashMap<String, String> status = new HashMap<>();
        status.put("status", "Ok!");
        return status;
    }

}
