package LZ09

import (
	"errors"
	"math"
	"problem"
	"solution"
)

var (
	numberOfVariables  int
	numberOfObjectives int
	pType              int
	lType              int
	dType              int
)

func lz09(nvar, nobj, ptype, ltype, dtype int) {
	numberOfVariables = nvar
	numberOfObjectives = nobj
	pType = ptype
	lType = ltype
	dType = dtype

}

func evaluate(solution *solution.Solution) {
	x_var := problem.GetReal(solution)

	var y_obj []float64 = make([]float64, numberOfObjectives, numberOfObjectives)

	if numberOfObjectives == 2 {
		if lType == 21 || lType == 22 || lType == 23 || lType == 24 || lType == 26 {
			var aa []float64 = make([]float64, 0)
			var bb []float64 = make([]float64, 0)

			for n := 1; n < numberOfVariables; n++ {
				if n%2 == 0 {
					aa = append(aa, psfunc2(x_var[n], x_var[0], n, 1))
				} else {
					bb = append(bb, psfunc2(x_var[n], x_var[0], n, 2))
				}
			}

			g := betafunction(aa)
			h := betafunction(bb)
			alpha := alphafunction(x_var)

			y_obj[0] = alpha[0] + h
			y_obj[1] = alpha[1] + g
		} else if lType == 25 {
			var aa []float64 = make([]float64, 0)
			var bb []float64 = make([]float64, 0)

			for n := 1; n < numberOfVariables; n++ {
				if n%3 == 0 {
					aa = append(aa, psfunc2(x_var[n], x_var[0], n, 1))
				} else if n%3 == 1 {
					bb = append(bb, psfunc2(x_var[n], x_var[0], n, 2))
				} else {
					c := psfunc2(x_var[n], x_var[0], n, 3)

					if n%2 == 0 {
						aa = append(aa, c)
					} else {
						bb = append(bb, c)
					}
				}
			}

			g := betafunction(aa)
			h := betafunction(bb)
			alpha := alphafunction(x_var)

			y_obj[0] = alpha[0] + h
			y_obj[1] = alpha[1] + g
		} else {
			errors.New("throw new IllegalStateException()")
		}
	} else if numberOfObjectives == 3 {
		if (lType == 31) || (lType == 32) {
			var aa []float64 = make([]float64, 0)
			var bb []float64 = make([]float64, 0)
			var cc []float64 = make([]float64, 0)

			for n := 2; n < numberOfVariables; n++ {
				a := psfunc3(x_var[n], x_var[0], x_var[1], n)

				if n%3 == 0 {
					aa = append(aa, a)
				} else if n%3 == 1 {
					bb = append(bb, a)
				} else {
					cc = append(cc, a)
				}
			}

			g := betafunction(aa)
			h := betafunction(bb)
			e := betafunction(cc)
			alpha := alphafunction(x_var)

			y_obj[0] = alpha[0] + h
			y_obj[1] = alpha[1] + g
			y_obj[2] = alpha[2] + e
		} else {
			errors.New("throw new IllegalStateException()")
		}
	} else {
		errors.New("throw new IllegalStateException()")
	}

	for i := 0; i < numberOfObjectives; i++ {
		solution.SetObjective(i, y_obj[i])
	}
}

/**
 * Validates the inputs, throwing an {@link IllegalArgumentException} if any
 * inputs are invalid.
 */
func validate() {
	if (numberOfObjectives < 2) || (numberOfObjectives > 3) {
		errors.New("invalid number of objectives")
	}

	if numberOfObjectives == 2 {
		if (pType < 21) || (pType > 24) {
			errors.New("invalid ptype")
		}
	} else if numberOfObjectives == 3 {
		if (pType < 31) || (pType > 34) {
			errors.New("invalid ptype")
		}
	}

	if (dType < 1) || (dType > 4) {
		errors.New("invalid dtype")
	}

	if numberOfObjectives == 2 {
		if (lType < 21) || (lType > 26) {
			errors.New("invalid ltype")
		}
	} else {
		if (lType < 31) || (lType > 32) {
			errors.New("invalid ltype")
		}
	}
}

/**
 * Controls the Pareto front shape.
 */
func alphafunction(x []float64) []float64 {
	var alpha []float64

	if numberOfObjectives == 2 {
		if pType == 21 {
			alpha = []float64{x[0], 1 - math.Sqrt(x[0])}
		} else if pType == 22 {
			alpha = []float64{x[0], 1 - x[0]*x[0]}
		} else if pType == 23 {
			alpha = []float64{x[0], 1 - math.Sqrt(x[0]) - x[0]*math.Sin(10*x[0]*x[0]*math.Pi)}
		} else if pType == 24 {
			alpha = []float64{x[0], 1 - x[0] - 0.05*math.Sin(4*math.Pi*x[0])}
		} else {
			errors.New("throw new IllegalStateException()")
		}
	} else if numberOfObjectives == 3 {
		if pType == 31 {
			alpha = []float64{math.Cos(x[0]*math.Pi/2) * math.Cos(x[1]*math.Pi/2), math.Cos(x[0]*math.Pi/2) * math.Sin(x[1]*math.Pi/2), math.Sin(x[0] * math.Pi / 2)}
		} else if pType == 32 {
			alpha = []float64{1 - math.Cos(x[0]*math.Pi/2)*math.Cos(x[1]*math.Pi/2), 1 - math.Cos(x[0]*math.Pi/2)*math.Sin(x[1]*math.Pi/2), 1 - math.Sin(x[0]*math.Pi/2)}
		} else if pType == 33 {
			alpha = []float64{x[0], x[1], 3 - (math.Sin(3*math.Pi*x[0]) + math.Sin(3*math.Pi*x[1])) - 2*(x[0]+x[1])}
		} else if pType == 34 {
			alpha = []float64{x[0] * x[1], x[0] * (1 - x[1]), (1 - x[0])}
		} else {
			errors.New("throw new IllegalStateException()")
		}
	} else {
		errors.New("throw new IllegalStateException()")
	}

	return alpha
}

/**
 * Controls the distance.
 */
func betafunction(x []float64) float64 {
	var beta float64

	dim := len(x)

	if dType == 1 {
		beta := 0.0

		for i := 0; i < dim; i++ {
			beta += x[i] * x[i]
		}

		return 2.0 * beta / float64(dim)
	} else if dType == 2 {
		beta := 0.0

		for i := 0; i < dim; i++ {
			beta += math.Sqrt(float64(i+1)) * x[i] * x[i]
		}

		beta = 2.0 * beta / float64(dim)
	} else if dType == 3 {
		sum := 0.0
		var xx float64

		for i := 0; i < dim; i++ {
			xx = 2 * x[i]
			sum += (xx*xx - math.Cos(4*math.Pi*xx) + 1)
		}

		beta = 2.0 * sum / float64(dim)
	} else if dType == 4 {
		sum := 0.0
		prod := 1.0
		var xx float64

		for i := 0; i < dim; i++ {
			xx = 2 * x[i]
			sum += xx * xx
			prod *= math.Cos(10 * math.Pi * xx / math.Sqrt(float64(i+1)))
		}

		beta = 2.0 * (sum - 2*prod + 2) / float64(dim)
	} else {
		errors.New("throw new IllegalStateException()")
	}

	return beta
}

/**
 * Controls the Pareto set shape for 2D instances.
 */
func psfunc2(x, t1 float64, dim, css int) float64 {
	var beta float64

	dim = dim + 1

	xy := 2 * (x - 0.5)

	if lType == 21 {
		beta = xy - math.Pow(t1, 0.5*(float64(numberOfVariables)+3*float64(dim)-8)/(float64(numberOfVariables)-2))
	} else if lType == 22 {
		theta := 6*math.Pi*t1 + float64(dim)*math.Pi/float64(numberOfVariables)
		beta = xy - math.Sin(theta)
	} else if lType == 23 {
		theta := 6*math.Pi*t1 + float64(dim)*math.Pi/float64(numberOfVariables)
		ra := 0.8 * t1
		if css == 1 {
			beta = xy - ra*math.Cos(theta)
		} else {
			beta = xy - ra*math.Sin(theta)
		}
	} else if lType == 24 {
		theta := 6*math.Pi*t1 + float64(dim)*math.Pi/float64(numberOfVariables)
		ra := 0.8 * t1
		if css == 1 {
			beta = xy - ra*math.Cos(theta/3)
		} else {
			beta = xy - ra*math.Sin(theta)
		}
	} else if lType == 25 {
		rho := 0.8
		phi := math.Pi * t1
		theta := 6*math.Pi*t1 + float64(dim)*math.Pi/float64(numberOfVariables)
		if css == 1 {
			beta = xy - rho*math.Sin(phi)*math.Sin(theta)
		} else if css == 2 {
			beta = xy - rho*math.Sin(phi)*math.Cos(theta)
		} else {
			beta = xy - rho*math.Cos(phi)
		}
	} else if lType == 26 {
		theta := 6*math.Pi*t1 + float64(dim)*math.Pi/float64(numberOfVariables)
		ra := 0.3 * t1 * (t1*math.Cos(4*theta) + 2)
		if css == 1 {
			beta = xy - ra*math.Cos(theta)
		} else {
			beta = xy - ra*math.Sin(theta)
		}
	} else {
		errors.New("throw new IllegalStateException()")
	}

	return beta
}

/**
 * Controls the Pareto set shape for 3D instances.
 */
func psfunc3(x, t1, t2 float64, dim int) float64 {
	var beta float64

	dim = dim + 1

	xy := 4 * (x - 0.5)

	if lType == 31 {
		rate := float64(dim) / float64(numberOfVariables)
		beta = xy - 4*(t1*t1*rate+t2*(1.0-rate)) + 2
	} else if lType == 32 {
		theta := 2*math.Pi*t1 + float64(dim)*math.Pi/float64(numberOfVariables)
		beta = xy - 2*t2*math.Sin(theta)
	} else {
		errors.New("throw new IllegalStateException()")
	}

	return beta
}
