package database

const createShema = `
CREATE TABLE IF NOT EXISTS info_vehicle
(
id varchar(256),
vehicleID varchar(256),
masterUserID varchar(256),
company varchar(256),
model varchar(256),
type varchar(256),
year int,
capacity int,
create_time timestamp,
modify_time timestamp
)
`

var insertRegistrationSchema = `
INSERT INTO info_vehicle(vehicleID, masterUserID, company, model, type, year, capacity,create_time, modify_time) VALUES($1,$2,$3,$4,$5,$6,$7, now(), now()) RETURNING id
`
