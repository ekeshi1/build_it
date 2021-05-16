create table plant_hires
(
    id         SERIAL PRIMARY KEY,
    plant_id        INTEGER NOT NULL,
    construction_site_id        INTEGER NOT NULL,
    supplier_id        INTEGER NOT NULL,
    site_engineer_id        INTEGER NOT NULL,
    plant_arrival_date date  NOT NULL,
    plant_return_date date  NOT NULL,
    plant_total_price FLOAT NOT NULL,
    plant_daily_price FLOAT NOT NULL,
    created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR NOT NULL
);

insert into plant_hires
        (plant_id,construction_site_id,supplier_id,site_engineer_id,plant_arrival_date,plant_return_date, plant_total_price, status)
values  
	(1,1,2,2,'2021-05-18','2021-05-20',123.5,'CREATED'),
        (2,3,1,1,'2021-05-19','2021-05-21',63.5,'CREATED'),
        (1,2,3,1,'2021-05-18','2021-05-20',123.5,'CREATED');
        
create table purchase_orders(
    id SERIAL PRIMARY KEY,
    plant_hire_id INTEGER NOT NULL,
    description VARCHAR ,
    created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    creator varchar NOT NULL,
    delivery_status varchar NOT NULL,
    status varchar NOT NULL,

    CONSTRAINT fk_plant
      FOREIGN KEY(plant_hire_id) 
	  REFERENCES plant_hires(id)
);

insert into purchase_orders
        (plant_hire_id,description,creator, delivery_status,status)
values  
	(1,'just a description1','BUILD_IT','CREATED','CREATED'),
        (2,'just a description2','BUILD_IT','CREATED','CREATED'),
        (3,'just a description3','BUILD_IT','CREATED','CREATED');
        
create table invoices(
    id SERIAL PRIMARY KEY,
    purchase_order_id INTEGER NOT NULL UNIQUE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    last_reminder_date date,
    payment_status varchar NOT NULL,
    
    CONSTRAINT fk_purchase_order
      FOREIGN KEY(purchase_order_id) 
	  REFERENCES purchase_orders(id)
);


insert into invoices(purchase_order_id, last_reminder_date, payment_status) values (1, '2021-06-15', 'CREATED'), (3, CURRENT_DATE, 'CREATED');
