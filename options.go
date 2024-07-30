package bitcask

import "os"

type IndexerType = int8

const (
	// BTree 索引
	BTree IndexerType = iota + 1

	// ART Adpative Radix Tree 自适应基数树索引
	ART

	// BPlusTree B+ 树索引，将索引存储到磁盘上
	BPlusTree
)

type Options struct {
	DirPath      string      // 数据库数据目录
	DataFileSize int64       // 数据文件大小阈值
	IndexType    IndexerType // 内存索引类型
	SyncWrites   bool        // 每次是否以持久化写入数据
}

// IteratorOptions 迭代器索引配置项
type IteratorOptions struct {
	// 遍历前缀为指定值的 key，默认为空
	Prefix []byte
	// 是否反向遍历，默认 false 是正向
	Reverse bool
}

// WriteBatchOptions 批量写配置项
type WriteBatchOptions struct {
	// 一个批次当中最大的数据量
	MaxBatchNum uint

	// 提交时是否 sync 持久化
	SyncWrites bool
}

var DefaultOptions = Options{
	DirPath:      os.TempDir(),
	DataFileSize: 256 * 1024 * 1024, // 256MB
	SyncWrites:   false,
	IndexType:    BTree,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}

var DefaultWriteBatchOptions = WriteBatchOptions{
	MaxBatchNum: 100000,
	SyncWrites:  true,
}
