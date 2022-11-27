package io.alpha.admin;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;
import xyz.erupt.core.annotation.EruptScan;

@SpringBootApplication
@ComponentScan({"xyz.erupt","io.alpha"}) // ↓ com.xxx要替换成实际需要扫描的代码包
@EntityScan({"xyz.erupt","io.alpha"})    // ↓ 例如DemoApplication所在的包为 com.example.demo
@EruptScan({"xyz.erupt","io.alpha"})
@EnableJpaRepositories({ "xyz.erupt","io.alpha"})// → 则：com.xxx → com.example.demo
//@EnableWebMvc
public class AlphaAdminApp {

    public static void main(String[] args) {
        SpringApplication.run(AlphaAdminApp.class);
    }
}
