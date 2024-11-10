
# Golyline

Golyline is a Go package for encoding and decoding geographical coordinates using the polyline algorithm. The polyline algorithm is a lossy compression technique commonly used in mapping applications, such as Google Maps, to efficiently store and transmit sequences of latitude/longitude pairs.

## Features

- **Encode**: Convert a sequence of [latitude, longitude] pairs into a polyline-encoded string.
- **Decode**: Convert a polyline-encoded string back into a sequence of [latitude, longitude] pairs.
- Efficient and optimized for performance, following best practices in Go.
- Comprehensive documentation and easy-to-use API.

## Installation

To use Golyline in your project, you can import it using Go modules:

```bash
go get github.com/amirhooshmand/golyline
```

Then, import it in your Go code:

```go
import "github.com/amirhooshmand/golyline"
```

## Usage

### Encoding Coordinates

You can encode a slice of [latitude, longitude] pairs into a polyline string:

```go
package main

import (
	"fmt"
	polyline "github.com/amirhooshmand/golyline"
)

func main() {
	points := [][]float64{
		{38.5, -120.2},
		{40.7, -120.95},
		{43.252, -126.453},
	}

	encoded := polyline.Encode(points)
	fmt.Println("Encoded polyline:", encoded)
}
```

### Decoding a Polyline

You can decode a polyline string back into a slice of [latitude, longitude] pairs:

```go
package main

import (
	"fmt"
	polyline "github.com/amirhooshmand/golyline"
)

func main() {
	polylineStr := "_p~iF~ps|U_ulLnnqC_mqNvxq`@"
	decoded := polyline.Decode(polylineStr)
	fmt.Println("Decoded coordinates:", decoded)
}
```

## API Reference

### `func Encode(points [][]float64) string`

- **Description**: Encodes a slice of [latitude, longitude] pairs into a polyline string.
- **Parameters**: `points` - A slice of [latitude, longitude] pairs as `[][]float64`.
- **Returns**: The polyline-encoded string.

### `func Decode(polyline string) [][]float64`

- **Description**: Decodes a polyline string back into a slice of [latitude, longitude] pairs.
- **Parameters**: `polyline` - The polyline-encoded string.
- **Returns**: A slice of [latitude, longitude] pairs as `[][]float64`.

### Internal Helper Functions

- **`encodeValue(value int) []byte`**: Encodes an integer using the polyline algorithm.
- **`decodeValue(polyline string, index *int) int`**: Decodes an integer from the polyline string.

## How It Works

The polyline algorithm encodes a series of geographical points into an ASCII string. The latitude and longitude values are multiplied by 1e5 and rounded to the nearest integer. The differences between each point and the previous one are encoded as a series of 5-bit chunks, with each chunk converted into a character.

For more details on the polyline algorithm, refer to the [Google Maps documentation](https://developers.google.com/maps/documentation/utilities/polylinealgorithm).

## Development

### Setting Up the Project

1. Clone the repository:
   ```bash
   git clone https://github.com/amirhooshmand/golyline.git
   cd golyline
   ```

2. Build the project:
   ```bash
   go build
   ```

3. Run tests:
   ```bash
   go test ./...
   ```

## Contributing

Contributions are welcome! If you'd like to improve the project or fix an issue, please:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m "Description of changes"`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Acknowledgments

- Inspired by the [Google Maps Polyline Algorithm](https://developers.google.com/maps/documentation/utilities/polylinealgorithm).

## Contact

For questions or feedback, please open an issue or reach out to [Amir Hooshmand](mailto:amir.hsmart@gmail.com).

---

Happy coding!
