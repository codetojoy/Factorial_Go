
package util

func GetChunk(index int, chunkSize int, max int) (int, int) {
    low := chunkSize * (index - 1)

    if low < 2 {
        low = 2
    }

    high := chunkSize * index

    if high > max {
        high = max
    }

    return low, high
}
