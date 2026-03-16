CREATE TABLE IF NOT EXISTS inside_contract(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    contract_id UUID REFERENCES production_contract(id),
    inside_contract_name VARCHAR(50) NOT NULL,
    inside_contract_price VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);