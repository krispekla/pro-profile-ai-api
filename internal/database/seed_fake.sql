DO $$
DECLARE table_name text;
BEGIN FOR table_name IN (
    SELECT tablename
    FROM pg_tables
    WHERE schemaname = 'public'
        AND tablename != 'schema_migrations'
) LOOP EXECUTE 'TRUNCATE TABLE ' || table_name || ' CASCADE';
END LOOP;
END $$;
-- INSERT INTO user_account (supa_id, name, email)
-- VALUES (
--         'supa_id_1',
--         'John Doe',
--         'outwork.trope.0x@icloud.com'
--     ),
--     (
--         'supa_id_2',
--         'Jane Smith',
--         'jane.smith@example.com'
--     );
-- Insert character section
INSERT INTO character (
        user_id,
        name,
        gender,
        hair_color,
        eye_color,
        ethnicity,
        age
    )
VALUES (
        (
            SELECT id
            FROM auth.users
            WHERE email = 'outwork.trope.0x@icloud.com'
        ),
        'John Doe',
        'male',
        'black',
        'brown',
        'light',
        'adult'
    ),
    (
        (
            SELECT id
            FROM auth.users
            WHERE email = 'outwork.trope.0x@icloud.com'
        ),
        'Jane Doe',
        'female',
        'blond',
        'blue',
        'medium',
        'adult'
    );
INSERT INTO character_img (character_id, img_url)
VALUES (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        2,
        'https://images.generated.photos/Sd_vJwUF7Bqg57Z0R_ekngFZ0EUOWtzpgmh-uwAzTA8/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/Mjk0OTk1LmpwZw.jpg'
    ),
    (
        2,
        'https://images.generated.photos/Sd_vJwUF7Bqg57Z0R_ekngFZ0EUOWtzpgmh-uwAzTA8/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/Mjk0OTk1LmpwZw.jpg'
    ),
    (
        2,
        'https://images.generated.photos/Sd_vJwUF7Bqg57Z0R_ekngFZ0EUOWtzpgmh-uwAzTA8/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/Mjk0OTk1LmpwZw.jpg'
    ),
    (
        2,
        'https://images.generated.photos/Sd_vJwUF7Bqg57Z0R_ekngFZ0EUOWtzpgmh-uwAzTA8/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/Mjk0OTk1LmpwZw.jpg'
    ),
    (
        2,
        'https://images.generated.photos/Sd_vJwUF7Bqg57Z0R_ekngFZ0EUOWtzpgmh-uwAzTA8/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/Mjk0OTk1LmpwZw.jpg'
    ),
    (
        2,
        'https://images.generated.photos/Sd_vJwUF7Bqg57Z0R_ekngFZ0EUOWtzpgmh-uwAzTA8/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/Mjk0OTk1LmpwZw.jpg'
    ),
    (
        2,
        'https://images.generated.photos/Sd_vJwUF7Bqg57Z0R_ekngFZ0EUOWtzpgmh-uwAzTA8/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/Mjk0OTk1LmpwZw.jpg'
    ),
    (
        2,
        'https://images.generated.photos/Sd_vJwUF7Bqg57Z0R_ekngFZ0EUOWtzpgmh-uwAzTA8/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/Mjk0OTk1LmpwZw.jpg'
    );
INSERT INTO character_model (character_id, model_url)
VALUES (1, 'http://example.com/john_doe_model.obj'),
    (2, 'http://example.com/jane_doe_model.svg');
-- Insert package section
INSERT INTO package (name, description, cover_img_url)
VALUES (
        'Linkedin package',
        'This LinkedIn package generates AI images tailored for professional networking. Enhance your online presence with high-quality, AI-generated images that align with LinkedIn platform aesthetics and user expectations.',
        'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f8/LinkedIn_icon_circle.svg/1024px-LinkedIn_icon_circle.svg.png'
    ),
    (
        'Christmas package',
        'This Christmas package generates AI images perfect for the holiday season. Create a festive atmosphere with high-quality, AI-generated images that capture the spirit of Christmas.',
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        'Instagram package',
        'This Instagram package generates AI images optimized for social media sharing. Boost your Instagram profile with high-quality, AI-generated images that resonate with the Instagram community.',
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    );
INSERT INTO package_price (package_id, amount, currency, stripe_product_id)
VALUES (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        20,
        'USD',
        'price_1OXMWLFSEa3MNRY9UD2qBbd0'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        20,
        'USD',
        'price_1OXMYbFSEa3MNRY94mt3xKhR'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        49,
        'USD',
        'price_1OXMXHFSEa3MNRY9y23Xz2DT'
    );
INSERT INTO coupon (package_id, code, amount, currency, percentage)
VALUES (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'COUPON10',
        10,
        'USD',
        NULL
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'COUPON20',
        20,
        'USD',
        NULL
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'COUPON30',
        NULL,
        NULL,
        30
    );
INSERT INTO package_example_img (package_id, img_url)
VALUES (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        ),
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        ),
        'https://images.thdstatic.com/productImages/abd4ce03-711d-428f-8bd6-17d7a6eda828/svn/home-accents-holiday-christmas-figurines-21sv22566-e1_600.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    ),
    (
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        ),
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg'
    );
INSERT INTO package_order (
        user_id,
        total_amount,
        currency,
        status,
        coupon_id
    )
VALUES (
        (
            SELECT id
            FROM auth.users
            WHERE email = 'outwork.trope.0x@icloud.com'
        ),
        10,
        'USD',
        'paid',
        (
            SELECT id
            FROM coupon
            WHERE code = 'COUPON10'
        )
    ),
    (
        (
            SELECT id
            FROM auth.users
            WHERE email = 'outwork.trope.0x@icloud.com'
        ),
        69,
        'USD',
        'paid',
        NULL
    ),
    (
        (
            SELECT id
            FROM auth.users
            WHERE email = 'outwork.trope.0x@icloud.com'
        ),
        20,
        'USD',
        'created',
        NULL
    ),
    (
        (
            SELECT id
            FROM auth.users
            WHERE email = 'outwork.trope.0x@icloud.com'
        ),
        20,
        'USD',
        'pending',
        NULL
    ),
    (
        (
            SELECT id
            FROM auth.users
            WHERE email = 'outwork.trope.0x@icloud.com'
        ),
        10,
        'USD',
        'cancelled',
        (
            SELECT id
            FROM coupon
            WHERE code = 'COUPON10'
        )
    ),
    (
        (
            SELECT id
            FROM auth.users
            WHERE email = 'outwork.trope.0x@icloud.com'
        ),
        20,
        'USD',
        'refunded',
        NULL
    );
INSERT INTO package_order_item (package_order_id, package_id)
VALUES (
        1,
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        )
    ),
    (
        2,
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        )
    ),
    (
        2,
        (
            SELECT id
            FROM package
            WHERE name = 'Instagram package'
        )
    ),
    (
        3,
        (
            SELECT id
            FROM package
            WHERE name = 'Christmas package'
        )
    ),
    (
        4,
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        )
    ),
    (
        5,
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        )
    ),
    (
        6,
        (
            SELECT id
            FROM package
            WHERE name = 'Linkedin package'
        )
    );
INSERT INTO generated_package (
        package_order_item_id,
        character_id,
        cover_img_url,
        status
    )
VALUES (
        2,
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg',
        'processing'
    ),
    (
        2,
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg',
        'generated'
    ),
    (
        3,
        1,
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg',
        'generated'
    ),
    (4, DEFAULT, NULL, DEFAULT),
    (5, DEFAULT, NULL, DEFAULT),
    (6, DEFAULT, NULL, DEFAULT);
INSERT INTO generated_package_img (generated_package_id, img_url, model_id)
VALUES (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg',
        1
    ),
    (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg',
        1
    ),
    (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg',
        1
    ),
    (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg',
        1
    ),
    (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg',
        1
    ),
    (
        1,
        'https://images.generated.photos/xvMS54uAAXqqavBGe82066GDFqKjkW0TgdA18kpNlEg/rs:fit:256:256/czM6Ly9pY29uczgu/Z3Bob3Rvcy1wcm9k/LnBob3Rvcy92M18w/OTE0MzI0LmpwZw.jpg',
        1
    ),
    (
        2,
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg',
        1
    ),
    (
        2,
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg',
        1
    ),
    (
        2,
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg',
        1
    ),
    (
        2,
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg',
        1
    ),
    (
        2,
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg',
        1
    ),
    (
        2,
        'https://images.ladbible.com/resize?type=jpeg&quality=70&width=648&fit=contain&gravity=null&url=https://s3-images.ladbible.com/s3/content/0c6495f0618693f2e69c848f47eba104.jpg',
        1
    );