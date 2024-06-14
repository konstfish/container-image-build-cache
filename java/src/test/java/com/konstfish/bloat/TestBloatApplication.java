package com.konstfish.bloat;

import org.springframework.boot.SpringApplication;

public class TestBloatApplication {

	public static void main(String[] args) {
		SpringApplication.from(BloatApplication::main).run(args);
	}

}
