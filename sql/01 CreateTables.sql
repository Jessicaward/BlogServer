create table Post (
    Title varchar not null,
    CreatedAt date not null
);

create table Tag (
    Key varchar not null,
    Name varchar not null
);

create table TagPost (
    PostId int not null,
    TagId int not null,
    FOREIGN KEY(PostId) REFERENCES Post(ROWID),
    FOREIGN KEY(TagId) REFERENCES Tag(ROWID)
);