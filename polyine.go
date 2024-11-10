package polyline

import (
	"math"
)

// Encode converts a slice of [latitude, longitude] pairs into a polyline-encoded string.
// The polyline encoding is a lossy compression algorithm that allows encoding a series of
// coordinates as a single string, commonly used in mapping applications.
//
// Parameters:
//
//	points: A slice of [latitude, longitude] pairs, where each pair is represented as
//	        a slice of two float64 values.
//
// Returns:
//
//	A string representing the polyline-encoded coordinates.
func Encode(points [][]float64) string {
	var result []byte
	prevLat, prevLng := 0, 0

	// Iterate through each point and encode the differences
	for _, point := range points {
		lat := int(math.Round(point[0] * 1e5))
		lng := int(math.Round(point[1] * 1e5))

		// Calculate the difference from the previous coordinates
		dLat := lat - prevLat
		dLng := lng - prevLng
		prevLat, prevLng = lat, lng

		// Encode the differences and append to the result
		result = append(result, encodeValue(dLat)...)
		result = append(result, encodeValue(dLng)...)
	}

	return string(result)
}

// Decode converts a polyline-encoded string back into a slice of [latitude, longitude] pairs.
//
// Parameters:
//
//	polyline: A string representing the polyline-encoded coordinates.
//
// Returns:
//
//	A slice of [latitude, longitude] pairs, where each pair is represented as
//	a slice of two float64 values.
func Decode(polyline string) [][]float64 {
	var (
		points   [][]float64
		index    int
		lat, lng int
	)

	// Iterate through the encoded string and decode each pair of values
	for index < len(polyline) {
		dLat := decodeValue(polyline, &index)
		dLng := decodeValue(polyline, &index)
		lat += dLat
		lng += dLng

		points = append(points, []float64{float64(lat) / 1e5, float64(lng) / 1e5})
	}

	return points
}

// encodeValue encodes a single integer value using the polyline encoding algorithm.
// The value is encoded as a series of 5-bit chunks, with each chunk converted to a character.
//
// Parameters:
//
//	value: The integer value to encode.
//
// Returns:
//
//	A slice of bytes representing the encoded value.
func encodeValue(value int) []byte {
	// Convert the value to a binary form, encoding the sign bit
	value = (value << 1) ^ (value >> 31)
	var encoded []byte

	// Encode each 5-bit chunk
	for value >= 0x20 {
		encoded = append(encoded, byte((0x20|(value&0x1f))+63))
		value >>= 5
	}
	encoded = append(encoded, byte(value+63))

	return encoded
}

// decodeValue decodes an integer value from the polyline-encoded string.
// The value is decoded from a series of 5-bit chunks, with each chunk read from the string.
//
// Parameters:
//
//	polyline: The polyline-encoded string.
//	index: A pointer to the current index in the string, which will be updated as characters are read.
//
// Returns:
//
//	The decoded integer value.
func decodeValue(polyline string, index *int) int {
	var result, shift int

	// Decode each 5-bit chunk and construct the integer value
	for {
		b := int(polyline[*index]) - 63
		*index++
		result |= (b & 0x1f) << shift
		shift += 5
		if b < 0x20 {
			break
		}
	}

	// Convert back from the binary form, decoding the sign bit
	return (result >> 1) ^ -(result & 1)
}
