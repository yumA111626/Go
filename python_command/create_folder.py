from pathlib import Path

# この関数は引数を元にファイルを作成する実行ディレクトリファイルを作成する(.go)
def create_GOfile(start_No , end_No , dir) :
    check_test_file = dir/"check.go" # 演習のファイルは作成フォルダの直下に作成する

    for num in range(start_No,end_No) : 
        folder_name = f"lesson_{num}"
        file_name = f"lesson{num}.go"

        create_folder = dir/folder_name
        create_file = create_folder/file_name
        
        try :
            Path(create_folder).mkdir(exist_ok=True)
            Path(create_file).touch()
            print(f"ファイルの作成に成功しました : {create_file}")
        except:
            print(f"ファイルの作成に失敗しました : {create_file}")

    try :
        Path(check_test_file).touch()
        print(f"ファイルの作成に成功しました : {check_test_file}")
    except:
        print(f"ファイルの作成に失敗しました : {check_test_file}")

if __name__ == "__main__":
    input("======= 本ファイルの使い方を説明します =======\n" \
    ".go ファイルを作成したいフォルダ名を入力する\n" \
    "番号はファイルの中身を弄って決定してください")

    target_dir = input("作成対象のフォルダ名を入力してください:")
    full_path = Path.cwd()/target_dir
    print(full_path)
    create_GOfile(start_No=39 , end_No=47 ,dir=full_path)