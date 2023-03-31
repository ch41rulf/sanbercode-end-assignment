-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
                       user_id INT PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       email VARCHAR(100),
                       phone_number VARCHAR(20),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
);

CREATE TABLE main_tasks (
                            task_id INT PRIMARY KEY,
                            title VARCHAR(100) NOT NULL,
                            description VARCHAR(255),
                            due_date DATE NOT NULL,
                            status VARCHAR(20) NOT NULL,
                            manager_id INT NOT NULL,
                            FOREIGN KEY (manager_id) REFERENCES users(user_id),
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
);

CREATE TABLE subtasks (
                          subtask_id INT PRIMARY KEY,
                          title VARCHAR(100) NOT NULL,
                          description VARCHAR(255),
                          status VARCHAR(20) NOT NULL,
                          main_task_id INT NOT NULL,
                          FOREIGN KEY (main_task_id) REFERENCES main_tasks(task_id),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
);

CREATE TABLE assignments (
                             assignment_id INT PRIMARY KEY,
                             main_task_id INT NOT NULL,
                             staff_id INT NOT NULL,
                             due_date DATE NOT NULL,
                             FOREIGN KEY (main_task_id) REFERENCES main_tasks(task_id),
                             FOREIGN KEY (staff_id) REFERENCES users(user_id),
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                             updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
);


-- +migrate StatementEnd