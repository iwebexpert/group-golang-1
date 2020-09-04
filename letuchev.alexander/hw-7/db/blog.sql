CREATE SCHEMA blog;

create table blog.t_post (
	nID         int not null,           
	sAbout      varchar(300),        
	sText       varchar(20000), 
	sLabels     varchar(10000),
	dtPublic 	timestamp     
);

create unique index t_post_pk on blog.t_post (nID);

CREATE SEQUENCE seq_post_id START 1;
