# gorountineTest

**本节我们用一个棋牌游戏服务器的例子来比较完整地展现Go语言并发编程的威力。因为我
们的重点不是网络编程，因此这个例子不会涉及网络层的细节。
另外，棋牌游戏通常由一组服务器协同以支持尽量多的同时在线玩家，但由于这种分布式
设计除了增加了局域网通信，从模型上与单服务器设计是一致的，或者说只相当于把多台服务
器的计算能力合并成逻辑上的一台单一服务器，所以本示例中我们只考虑单服务器、单进程的设
计方法。
首先我们来分析这个项目的详细需求。作为一个棋牌游戏，需要支持玩家进行下面的基本
操作：
 登录游戏
 查看房间列表
 创建房间
 加入房间
 进行游戏
 房间内聊天
 游戏完成，退出房间
 退出登录**

**棋牌游戏的特点在于房间与房间之间具备良好的隔离性，这也是最能够体现并行编程威力的
地方。因为goroutine可创建的个数不受系统资源的限制，原则上一台服务器可以创建上百万个
goroutine，也就是可能可以支撑上百万个房间。当然，考虑到每个房间都需要耗费计算和内存资
源，实际上不可能达到这么高的数字，但我们可以预测与使用系统线程和系统进程来对应一个房
间相比，显然使用goroutine可以支持得多很多。
接下来我们开始进行系统设计。先简化登录流程：用户只需要输入用户名就可以直接登录，
无需验证过程。因此，对于用户管理，就是一个会话的管理流程。每个玩家对应的信息如下：
 用户唯一ID 
 用户名，用于显示
 玩家等级
 经验值
实际的游戏设计当然要比这个复杂得多，比如还有社交关系、道具和技能等。鉴于这些细节
并不影响架构，这里我们都一并略去。
总体上，我们可以将该示例划分为以下子系统：
 玩家会话管理系统，用于管理每一位登录的玩家，包括玩家信息和玩家状态
 大厅管理
 房间管理系统，创建、管理和销毁每一个房间
 游戏会话管理系统，管理房间内的所有动作，包括游戏进程和房间内聊天
 聊天管理系统，用于接收管理员的广播信息
为了避免贴出太多源代码，这里我们只实现了最基础的会话管理系统和聊天管理系统。因为
它们足以展示以下的技术问题：
 goroutine生命周期管理
 goroutine之间的通信
 共享资源访问控制
其他子系统所使用的技术与我们实现的代码是完全一致的，只不过需要便携不同的业务代
码。因此，相信即使我们没有实现所有子模块，如果读者有兴趣的话，要将其完整实现也并非
难事。
我们的目录结构如下：**





```
<cgss> 
├─<src> 
 
    ├─<cg> 
 
        ├─center.go 
 
        ├─centerclient.go 
 
        ├─player.go 
 
     ├─<ipc> 
 
         ├─server.go 
 
         ├─client.go 
 
         ├─ipc_test.go 
 
     ├─cgss.go 

```

