/*
[機能]
1. 全角アルファベット,数字を半角に変換する。
2. ローマ数字を半角数字に変換
 */

drop function if exists translate_case(text);
create function translate_case(
    p_src text
) returns text as $$
declare
    text_result text := upper(p_src);
    arr_roman_nums text[] := array['I', 'Ⅱ', 'Ⅲ', 'Ⅳ', 'Ⅴ', 'Ⅵ', 'Ⅶ', 'Ⅷ', 'Ⅸ', 'Ⅹ'];
    arr_hankaku_nums text[] := array['1', '2', '3', '4', '5', '6', '7', '8', '9', '10'];
    text_zenkaku_nums text := '０１２３４５６７８９';
    text_hankaku_nums text := '0123456789';
    text_zenkaku_upper_alphabets text := 'ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ';
    text_zenkaku_lower_alphabets text := 'ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ';
    text_hankaku_upper_alphabets text := 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
begin
    if p_src is null or p_src = '' then
        return p_src;
    end if;

    for i in 1..array_length(arr_roman_nums, 1) loop
        text_result := replace(text_result, arr_roman_nums[i], arr_hankaku_nums[i]);
    end loop;

    text_result := translate(
        text_result
        , text_zenkaku_upper_alphabets || text_zenkaku_lower_alphabets || text_zenkaku_nums
        , text_hankaku_upper_alphabets || text_hankaku_upper_alphabets || text_hankaku_nums
    );

    return text_result;
end;
$$ language plpgsql immutable;