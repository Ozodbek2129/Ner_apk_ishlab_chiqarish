CREATE TABLE IF NOT EXISTS production_contract(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    contract_name VARCHAR(50) UNIQUE NOT NULL,
    contract_number VARCHAR(50) NOT NULL,
    contract_deadline VARCHAR(50) NOT NULL,
    contract_file TEXT DEFAULT 'image',
    responsible_person VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);