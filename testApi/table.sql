CREATE TABLE public.users (
	user_id uuid NOT NULL,
	first_name varchar(255) NOT NULL,
	last_name varchar(255) NOT NULL,
	phone_number varchar(20) NULL,
	address varchar(255) NULL,
	pin varchar(10) NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE public.transactions (
	transaction_id uuid NOT NULL,
	user_id uuid NOT NULL,
	status varchar(50) NOT NULL,
	transaction_type varchar(10) NOT NULL,
	amount numeric(10, 2) NOT NULL,
	remarks text NULL,
	balance_before numeric(10, 2) DEFAULT 0 NOT NULL,
	balance_after numeric(10, 2) DEFAULT 0 NOT NULL,
	created_date timestamp DEFAULT now() NULL,
	CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id)
);