|      name       | description                              |
| :-------------: | :--------------------------------------- |
|      ping       | ping                                     |
|       now       | 伺服器現在時間                           |
|  countdown-R0   | R0 的倒數計時                            |
|     card-R0     | 所有關於 R0 的字卡更新，用於 /card/admin |
| card-current-R0 | R0 現在字卡的更新，用於 /card?room=R0    |
| card-id-2d8a5e  | 某個字卡的更新，用於 /card?id=2d8a5e     |

# TODO

-   [ ] 把 `Name` 改成 `Topic`，也就是說一條訊息可以包含好幾個 Topic，節省流量
-   [x] ~~當 now 更新時，要檢查要不要重送 card-current-R0~~ 直接在收到 MsgNow 時順便更新所有現在字卡
-   [x] ~~當某個字卡因為更新時間不再是目前字卡時，需要通知 card-current-R0~~ 修正 UpdateCardInRoom
