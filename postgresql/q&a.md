# Q
migrationをした段階ではtableは一つずつじゃないですか。その後にuserを追加したりdiaryを追加する中で、多対多の関係になります。この時、tableは一個しかないのに何で複数のuserとdiaryができるんですか？

# A
create tableは設計図であり、userはゼロ人
insertでprimary keyが重なった時に複数のuserやdiaryが作成される


# Q
これを考えると中間テーブルって意味あるんですか？無くても全部できますよね？

# A
今回は一人のuserに複数のdiary,一つのdiaryに一人のuserだから一対多
中間テーブルが必要になる時は多対多の時
一つのdiaryに複数のtag,一つのtagに複数の日記の時などに使用
