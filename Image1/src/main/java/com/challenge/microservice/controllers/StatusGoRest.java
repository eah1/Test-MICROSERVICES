package com.challenge.microservice.controllers;

public class StatusGoRest {
    private String status;

    public StatusGoRest(String status) {
        this.status = status;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

}
