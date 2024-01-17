create trigger "user_registration_webhook"
after
insert on "auth"."users" for each row execute function "supabase_functions"."http_request"(
        'https://gj0gc7x2-3002.euw.devtunnels.ms/api/user/webhook',
        'POST',
        '{"Content-Type":"application/json"}',
        '{}',
        '1000'
    );