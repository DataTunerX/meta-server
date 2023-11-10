### LLM Dataset CRD:

---

#### **ApiVersion:** 
core.datatunerx.io/v1beta1

#### **Kind:** 
Dataset

#### **Metadata:** 
- **Name:** 数据集名称

#### **Spec:**  
- **DatasetMetadata:**  
  - **Languages:** 该数据集使用的语言  
  - **Tags:** 该数据集自定义标签  
  - **License:** 该数据集许可信息  
  - **Size:** 数据集的样本规模  
  - **Task:** 
    - **Name:** 该任务的名称
    - **SubTasks:**  
      - **Name:** 子任务名称
      - **Name:** 子任务名称 1
  - **DatasetInfo:**  
    - **Subsets:**  
      - **Name:** 子数据集名称默认Default  
      - **Splits:**  
        - **Train:**  
          - **File:** 训练数据集数据包地址  
        - **Test:**  
          - **File:** 测试数据集数据包地址  
        - **Val:**  
          - **File:** 验证数据集数据包地址  
    - **Features:**  
      - **Name:** 微调数据字段
      - **MapTo:** 数据集特征值 
      - **DataType:** 数据类型
  - **Plugin:** 
    - **LoadPlugin:** true 是否使用的是插件  
    - **Name:** My plugin 1  
    - **Parameters:** 数据集插件参数，如 "{'params1': 'value1', 'params2': 'value2}"
- **DatasetCard:** 数据集的 Readme 信息  
  - **DatasetCardRef:** 一个 configmap 挂载的 Readme 文件  
- **DatasetFiles:**  
  - **Source:** 数据集文件源  

#### **Status:**  
- **State:** READY 数据集可用，UNREADY 数据集不可用

---

### FinetuneJob CRD:

---

#### **ApiVersion:** 
finetune.datatunerx.io/v1beta1

#### **Kind:** 
FinetuneJob

#### **Metadata:** 
- **Name:** 流水线名称  

#### **Spec:** 
- **FineTune:**  
  - **LLM:** 使用的大模型 CR  
  - **Dataset:** 使用的数据集 CR  
  - **Hyperparameter:** 使用的超参 CR  
  - **Resource:**  
    - **Limits:**  
      - **Cpu:** cpu 限制  
      - **Memory:** 内存限制  
      - **Gpu:** gpu 限制
- **ServeConfig:**  
  - **NodeSelector:** 设置 Ray 服务部署的节点  
  - **Tolerations:** 设置 Ray 服务容忍  
- **ScoringPlugin:**
  - **Name:** 使用的scoringplugin cr 名称

#### **Status:**  
- **State:** INIT/FAILED/SUCCESSFUL/BUILDIMAGE/FINETUNE/SERVE  
- **FinetuneState:**  
  - **State:** Finetune 的状态  
  - **Stats:** Finetune 的统计信息  
  - **CheckPoint:** Finetune 检查点  
- **Stats:** 工作流运行统计信息，开始时间与结束时间  
- **Result:**  
  - **ModelExportResult:** 模型导出结果 true  
  - **Dashboard:** Ray 服务的 dashboard
  - **Serve:** 模型服务地址
  - **Image:** 镜像

---

### FinetuneExperiment CRD:

---

#### **ApiVersion:** 
finetune.datatunerx.io/v1beta1

#### **Kind:**
FinetuneExperiment

#### **Metadata:** 
- **Name:** 微调工作名称  

#### **Spec:** 
- **FinetuneJobs:** 多个 FinetuneJob CR 的 Spec
- **ScoringConfig:** 
  - **Name:** My-scoring scoring cr 名称  

#### **Status:** 
- **BestVersion:** 
  - **Score:** 得分  
  - **Image:** 镜像  
  - **LLM:** 使用的大模型  
  - **Hyperparameter:** 使用的超参数  
  - **Dataset:** 使用的数据集  

- **JobStates:** 该 job 下的每条流水线信息
  - **FinetuneJob-0:**
    - **State:** RUNNING  
    - **Stats:** 2006-10-22  
    - **DashBoard:** xxxxx  
    - **CheckPoint:** v1.0  
    - **Score:** 0.1
  - **FinetuneJob-1:** (and so on...)

- **State:** Processing/Success/Failed 该 Experiment 状态

---

### LLM CRD

---

#### **ApiVersion:** 
core.datatunerx.io/v1beta1

#### **Kind:** 
LLM

#### **Metadata:** 
- **Name:** 模型名称  

#### **Spec:** 
- **LLMMetadata:**  
  - **Name:** 模型名称
  - **Domain:** 模型所涉及的领域
  - **Tasks:** 模型对应的的任务类型  
  - **Frameworks:** 模型架构  
  - **Languages:** 模型支持的自然语言列表  
  - **License:** 模型许可证信息  
  - **Tags:** 自定义标签，用于检索模型  
  - **Datasets:** 模型使用到的数据集  
  - **BaseLLM:** 模型的基础模型  
- **LLMImageConfig:**  
  - **Image:** 镜像地址  
  - **path:** 源文件地址  
- **ComputeInfrastructure:**  
  - **Hardware:** 硬件  
  - **VRam:** 显存  
  - **Cpu:** cpu  
  - **Memory:** 内存大小  
- **Version:** 模型的版本  
- **LLMCard:** 模型的 Readme 信息  
- **LLMCardRef:** 一个 configmap 挂载的 Readme 文件  
- **LLMFiles:**  
  - **Source:** 模型源文件引用地址  

---

### Hyperparameter CRD

---

#### **ApiVersion:** 
core.datatunerx.io/v1beta1

#### **Kind:** 
Hyperparameter

#### **Metadata:** 
- **Name:** 超参 CR 名称  超参数的名称标识

#### **Spec:** 

- **Objective:**  目标定义部分
  - **Type:** 微调类型  定义微调的具体类型

- **Parameters：**  超参数列表部分
  - **Scheduler：**  调度器类型
  - **Optimizer：**  优化器类型
  - **Int4/8:**  整数参数，可能是量化的一部分
  - **LoRA_R:**  LoRA的R参数
  - **LoRA_Alpha:**  LoRA的Alpha参数
  - **LoRA_Dropout:**  LoRA的Dropout参数
  - **LearningRate:**  学习速率
  - **Epochs:**  总的训练周期数
  - **BlockSize:**  块的大小
  - **BatchSize:**  每批的样本数量
  - **WarmupRatio:**  预热比率
  - **WeightDecay:**  权重衰减
  - **GradAccSteps:**  梯度积累步骤
  - **TrainerType:**  训练器类型
  - **PEFT:** bool  PEFT标志，可能是一个布尔值标志
  - **FP16:** bool  是否使用16位浮点数训练

---

### DataPlugin CRD

---

#### **ApiVersion:** 
extension.datatunerx.io/v1alpha

#### **Kind:** 
DataPlugin

#### **Metadata:** 
- **Name:** my-DataPlugin

#### **Spec:** 
- **DatasetClass：** 插件类型
- **Version:** 插件版本
- **Provider:** 提供商
- **Parameters:** 插件参数部分，如 "{'params1': 'value1', 'params2': 'value2}"

---

### ScoringPlugin CRD

---

#### **ApiVersion:** 
extension.datatunerx.io/v1alpha

#### **Kind:** 
ScoringPlugin

#### **Metadata:** 
- **Name:** my-ScoringPlugin

#### **Spec:** 
- **Version:** 插件版本
- **ScoringClass：** 打分插件类型
- **Metrics：** 插件支持的评价指标
  - Accu
  - Pre
  - F1
- **Parameters:** 插件参数部分，如 "{'params1': 'value1', 'params2': 'value2}"

---

### Scoring CRD

---

#### **ApiVersion:** 
extension.datatunerx.io/v1alpha

#### **Kind:** 
Scoring

#### **Metadata:** 
- **Name:** Scoring1

#### **Spec:** 
- **Plugins:** 插件配置部分
  - **LoadPlugin:** 是否使用插件
  - **Name:** 插件名称
  - **Parameters:** 插件参数部分, 如  "{'params1': 'value1', 'params2': 'value2}"
- **Questions:** 问题列表
  - **Question:** 问题描述
    **Reference：** 标准答案
  - **Question:** 另一个问题描述
    **Reference：** 对应的标准答案
- **Address**：进行打分的调用地址

---

### Finetune CRD

---

#### **ApiVersion:** 
finetune.daocloud.io/v1beta1

#### **Kind:** 
Finetune

#### **Metadata:** 
- **Name:** finetune1

#### **Spec:** 
- **Hyperparameter:** 使用的超参数 CR 名称
- **Dataset:** 使用的数据集 CR 名称
- **Llm:** 使用的大模型 CR 名称
- **Node:** 节点数量
- **NodeSelector:** 节点选择
- **Resource:** 资源配置部分
  - **Cpus:** CPU 数量
  - **Memory:** 内存大小
  - **Gpus:** GPU 数量
- **Image** 服务镜像
  - **Name** 镜像名称
  - **ImagePullPolicy** 镜像拉取策略
  - **Path** 镜像中模型文件的地址

#### **Status:** 
- **State:** 运行状态 (例如：init、pending、running、failed、successful)

---

### LLMCheckpoint CRD

---

#### **ApiVersion:** 
core.daocloud.io/v1beta1

#### **Kind:** 
LLMCheckpoint

#### **Metadata:** 
- **Name:** Checkpoint1

#### **Spec:** 
- **Hyperparameter:** 使用的超参数对象的spec
- **Dataset:** 使用的数据集对象的spec
- **Llm:** 使用的大模型对象的spec
- **Image:** 镜像配置部分
  - **Name:** 镜像名称
  - **ImagePullPolicy:** 拉取策略
  - **Path:** 大模型路径
- **Checkpoint:** checkpoint 地址
- **CheckpointImage:** checkpoint镜像配置部分
  - **Name:** 镜像名称
  - **ImagePullPolicy:** 拉取策略
  - **LLMPath:** 大模型路径
  - **CheckPointPath：** checkpoint 地址

---




