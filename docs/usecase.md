```mermaid

%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#ffcc00', 'edgeLabelBackground':'#ffffff', 'tertiaryColor': '#ffecb3'}}}%%

graph LR
    actor_User[一般ユーザー]
    actor_Admin[管理ユーザー]

    subgraph 一般機能
        UC_Login[ログイン/ログアウト]
        UC_RegisterSkill[スキル登録]
        UC_SearchUser[ユーザー検索]
        UC_ViewProfile[プロフィール閲覧]
        UC_EditProfile[プロフィール編集]
    end

    subgraph 管理ユーザー機能
        UC_RegisterUser[ユーザー登録]
        UC_CreateEditSkill[スキル作成/編集]
    end

    actor_User --> UC_Login
    actor_User --> UC_RegisterSkill
    actor_User --> UC_SearchUser
    actor_User --> UC_ViewProfile
    actor_User --> UC_EditProfile

    actor_Admin --> UC_Login
    actor_Admin --> UC_RegisterSkill
    actor_Admin --> UC_SearchUser
    actor_Admin --> UC_ViewProfile
    actor_Admin --> UC_EditProfile
    actor_Admin --> UC_RegisterUser
    actor_Admin --> UC_CreateEditSkill



```
