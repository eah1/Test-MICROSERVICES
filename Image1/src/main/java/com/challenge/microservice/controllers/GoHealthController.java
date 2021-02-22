package com.challenge.microservice.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import java.util.HashMap;

@RestController
public class GoHealthController {

    @GetMapping(value = "health/go")
    public StatusGoRest handlerGoHealth() {

        RestTemplate restTemplate = new RestTemplate();
        StatusGoRest status = restTemplate.getForObject("http://image2_micro2_1:5010/health", StatusGoRest.class);

        return status;
    }



}
