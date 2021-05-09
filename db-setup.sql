create table plantHire
(
    id         SERIAL PRIMARY KEY,
    plantId        INTEGER NOT NULL,
    constructionSiteId        INTEGER NOT NULL,
    supplier        INTEGER NOT NULL,
    siteEngineerId        INTEGER NOT NULL,
    plantArrivalDate date  NOT NULL,
    plantReturnDate date  NOT NULL,
    plantTotalPrice FLOAT NOT NULL,
    createdAt timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR NOT NULL
);

insert into plantHire
        (plantId,constructionSiteId,supplier,siteEngineerId,plantArrivalDate,plantReturnDate, plantTotalPrice, status)
values  
	(1,1,2,2,'2021-05-18','2021-05-20',123.5,'CREATED'),
        (2,3,1,1,'2021-05-19','2021-05-21',63.5,'CREATED'),
        (1,2,3,1,'2021-05-18','2021-05-20',123.5,'CREATED');
        
create table purchaseOrder(
    id SERIAL PRIMARY KEY,
    plantHireId INTEGER NOT NULL,
    description VARCHAR ,
    createdAt timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    creator varchar NOT NULL,
    deliveryAddress varchar NOT NULL,
    deliveryStatus varchar NOT NULL,
  	status varchar NOT NULL,

    CONSTRAINT fk_plant
      FOREIGN KEY(plantHireId) 
	  REFERENCES plantHire(id)
);

insert into purchaseOrder
        (plantHireId,description,creator,deliveryAddress, deliveryStatus,status)
values  
	(1,'just a description1','BUILD_IT','Address1','CREATED','CREATED'),
        (2,'just a description2','BUILD_IT','Address2','CREATED','CREATED'),
        (3,'just a description3','BUILD_IT','Address3','CREATED','CREATED');
        
create table invoice(
    id SERIAL PRIMARY KEY,
    purchaseOrderId INTEGER NOT NULL,
    createdDate timestamp DEFAULT CURRENT_TIMESTAMP,
    updatedDate timestamp DEFAULT CURRENT_TIMESTAMP,
    lastReminderDate date,
    paymentStatus varchar NOT NULL,
    
    CONSTRAINT fk_purchase_order
      FOREIGN KEY(purchaseOrderId) 
	  REFERENCES purchaseOrder(id)
);


insert into invoice(purchaseOrderId, lastReminderDate, paymentStatus) values (1, '2021-06-15', 'CREATED'), (3, CURRENT_DATE, 'CREATED');
