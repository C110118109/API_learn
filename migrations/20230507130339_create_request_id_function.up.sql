create function add_request_code()
create function add_request_code()
returns text as
$$
declare
    old_id text :=(select  request_code  from requests order by request_code desc limit 1);
    id_number varchar(4) :='0001';
    -- datetime char(4) :=substring(cast(now() as varchar),1,4);
    new_id text ;
    num integer;
begin
    if old_id is null then
        new_id:='P'||id_number;
        return new_id;
    end if;
    
   
    num :=cast(right(old_id,4) as integer)+1;
    id_number:=
    case
        when num<10 then '000'||num
        when num<100 then '00'||num
		when num<1000 then '0'||num
        when num<10000 then cast(num as text)
    end;
    
    
    new_id:='P'||id_number;
    return new_id;
end; 
$$
language 'plpgsql';