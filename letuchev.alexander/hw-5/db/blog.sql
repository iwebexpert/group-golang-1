CREATE SCHEMA blog;

CREATE TABLE blog.t_post (
	nid int4 NOT NULL,
	sabout varchar(300) NULL,
	stext varchar(20000) NULL,
	slabels varchar(10000) NULL,
	dtpublic timestamptz NULL
);
CREATE UNIQUE INDEX t_post_pk ON blog.t_post USING btree (nid);

CREATE SEQUENCE seq_post_id START 1;
