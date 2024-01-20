create trigger "user_registration_webhook"
after
insert on "auth"."users" for each row execute function "supabase_functions"."http_request"(
        'http://localhost:3002/api/user/webhook',
        'POST',
        '{"Content-Type":"application/json"}',
        '{}',
        '1000'
    );