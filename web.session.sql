INSERT INTO user_profile (
    id,
    name,
    follow_count,
    follower_count,
    avater,
    background_image,
    signature,
    total_favorited,
    work_count,
    favorite_count,
    created_at,
    updated_at,
    deleted_at
  )
VALUES (
    '1',
    'test01',
    0,
    0,
    'http://someurl.com',
    'http://someurl.com',
    '这个用户很懒,什么也没留下',
    '0',
    1,
    0,
    '2003-08-04 15:50:30',
    '2003-08-04 15:50:30',
    NULL
  );
INSERT INTO user_credentials (
    user_id,
    username,
    password,
    created_at,
    updated_at,
    deleted_at
  )
VALUES (
    '1',
    'test01',
    '123456',
    '2003-08-04 15:50:30',
    '2003-08-04 15:50:30',
    NULL
  );

INSERT INTO video_info (
    id,
    user_id,
    play_url,
    cover_url,
    favorite_count,
    comment_count,
    title,
    created_at,
    updated_at,
    deleted_at
  )
VALUES (
    '1',
    '1',
    'http://someurl.com',
    '3万',
    '0',
    '0',
    'title',
    '2003-08-04 15:50:30',
    '2003-08-04 15:50:30',
    NULL
  );