    CREATE TABLE groups (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        name TEXT NOT NULL,
        master VARCHAR(50) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        deleted_at BIGINT DEFAULT 0
    );

    CREATE TABLE workers (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        group_id UUID REFERENCES groups(id) ON DELETE CASCADE,
        name TEXT NOT NULL,
        image TEXT DEFAULT 'image',    
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        deleted_at BIGINT DEFAULT 0
    );

    CREATE TABLE attendance (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        worker_id UUID REFERENCES workers(id) ON DELETE CASCADE,
        situation VARCHAR(70) NOT NULL,
        reason VARCHAR(50) NOT NULL,
        is_present BOOLEAN NOT NULL,
        work_date DATE NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        deleted_at BIGINT DEFAULT 0
    );

    CREATE TABLE tasks (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        worker_id UUID REFERENCES workers(id) ON DELETE CASCADE,
        task TEXT NOT NULL,
        deadline VARCHAR(50) NOT NULL,
        start_time TIMESTAMP NOT NULL,
        end_time TIMESTAMP NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        deleted_at BIGINT DEFAULT 0
    );

    CREATE INDEX idx_tasks_worker ON tasks(worker_id);

    CREATE INDEX idx_attendance ON attendance(worker_id);  