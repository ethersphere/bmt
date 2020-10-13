package zerohash

var lookup [][]byte
const depth

func init() {
	depth := calculateDepthFor(segmentCount)
	zerohashes := make([][]byte, depth+1)
	zeros := make([]byte, segmentSize)
	zerohashes[0] = zeros
	h := hasher()
	for i := 1; i < depth+1; i++ {
		zeros = doSum(h, nil, zeros, zeros)
		zerohashes[i] = zeros
	}

}


// calculateDepthFor calculates the depth (number of levels) in the BMT tree.
func calculateDepthFor(n int) (d int) {
	c := 2
	for ; c < n; c *= 2 {
		d++
	}
	return d + 1
}


