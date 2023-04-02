-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
                       user_id SERIAL PRIMARY KEY,
                       username VARCHAR(50),
                       password VARCHAR(255),
                       email VARCHAR(100),
                       phone_number VARCHAR(20),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);


CREATE TABLE main_tasks (
                            task_id SERIAL PRIMARY KEY,
                            title VARCHAR(100) ,
                            description VARCHAR(255),
                            due_date DATE ,
                            status VARCHAR(20) ,
                            manager_id INT ,
                            FOREIGN KEY (manager_id) REFERENCES users(user_id),
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE subtasks (
                          subtask_id SERIAL PRIMARY KEY,
                          title VARCHAR(100) ,
                          description VARCHAR(255),
                          status VARCHAR(20) ,
                          main_task_id INT ,
                          FOREIGN KEY (main_task_id) REFERENCES main_tasks(task_id),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE assignments (
                             assignment_id SERIAL PRIMARY KEY,
                             main_task_id INT ,
                             staff_id INT ,
                             due_date DATE ,
                             FOREIGN KEY (main_task_id) REFERENCES main_tasks(task_id),
                             FOREIGN KEY (staff_id) REFERENCES users(user_id),
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                             updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +migrate StatementEnd
