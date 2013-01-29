/*
 * Easing - https://github.com/paweldubiel/easing
 *
 * A port of Robert Penner's easing equations to GO 
 * 
 * TERMS OF USE - Easing
 * 
 * Open source under the BSD License. 
 * 
 * Copyright © 2013 Pawel Dubiel
 * All rights reserved.
 * 
 * Redistribution and use in source and binary forms, with or without modification, 
 * are permitted provided that the following conditions are met:
 * 
 * Redistributions of source code must retain the above copyright notice, this list of 
 * conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright notice, this list 
 * of conditions and the following disclaimer in the documentation and/or other materials 
 * provided with the distribution.
 * 
 * Neither the name of the author nor the names of contributors may be used to endorse 
 * or promote products derived from this software without specific prior written permission.
 * 
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY 
 * EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 *  COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 *  EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE
 *  GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED 
 * AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 *  NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED 
 * OF THE POSSIBILITY OF SUCH DAMAGE. 
 */

package easing

import (
	"math"
)

//t: current time, b: begInnIng value, c: change In value, d: duration

func Linear(t, b, c, d float64) float64 {
	return c*t/d + b
}

func EaseInQuad(t, b, c, d float64) float64 {
	t /= d
	return c*t*t + b
}

func EaseOutQuad(t, b, c, d float64) float64 {
	t /= d
	return -c*t*(t-2) + b
}

func EaseInOutQuad(t, b, c, d float64) float64 {
	t /= d / 2
	if t < 1 {
		return c/2*t*t + b
	}
	t--
	return -c/2*(t*(t-2)-1) + b
}

func EaseInCubic(t, b, c, d float64) float64 {
	t /= d
	return c*t*t*t + b
}

func easeOutCubic(t, b, c, d float64) float64 {
	t /= d
	t--
	return c*(t*t*t+1) + b
}

func EaseInOutCubic(t, b, c, d float64) float64 {
	t /= d / 2
	if t < 1 {
		return c/2*t*t*t + b
	}
	t -= 2
	return c/2*(t*t*t+2) + b
}

func EaseInQuart(t, b, c, d float64) float64 {
	t /= d
	return c*t*t*t*t + b
}

func EaseOutQuart(t, b, c, d float64) float64 {
	t /= d
	t--
	return -c*(t*t*t*t-1) + b
}

func EaseInOutQuart(t, b, c, d float64) float64 {
	t /= d / 2
	if t < 1 {
		return c/2*t*t*t*t + b
	}
	t -= 2
	return -c/2*(t*t*t*t-2) + b
}

func EaseInQuint(t, b, c, d float64) float64 {
	t /= d
	return c*t*t*t*t*t + b
}

func EaseOutQuint(t, b, c, d float64) float64 {
	t /= d
	t--
	return c*(t*t*t*t*t+1) + b
}

func EaseInOutQuint(t, b, c, d float64) float64 {
	t /= d / 2
	if t < 1 {
		return c/2*t*t*t*t*t + b
	}
	t -= 2
	return c/2*(t*t*t*t*t+2) + b
}

func EaseInSine(t, b, c, d float64) float64 {
	return -c*math.Cos(t/d*(math.Pi/2)) + c + b
}

func EaseOutSine(t, b, c, d float64) float64 {
	return c*math.Sin(t/d*(math.Pi/2)) + b
}

func EaseInOutSine(t, b, c, d float64) float64 {
	return -c/2*(math.Cos(math.Pi*t/d)-1) + b
}

func EaseInExpo(t, b, c, d float64) float64 {
	return c*math.Pow(2, 10*(t/d-1)) + b
}

func EaseOutExpo(t, b, c, d float64) float64 {
	return c*(-math.Pow(2, -10*t/d)+1) + b
}

func EaseInOutExpo(t, b, c, d float64) float64 {
	t /= d / 2
	if t < 1 {
		return c/2*math.Pow(2, 10*(t-1)) + b
	}
	t--
	return c/2*(-math.Pow(2, -10*t)+2) + b
}

func EaseInCirc(t, b, c, d float64) float64 {
	t /= d
	return -c*(math.Sqrt(1-t*t)-1) + b
}

func EaseOutCirc(t, b, c, d float64) float64 {
	t /= d
	t--
	return c*math.Sqrt(1-t*t) + b
}

func EaseInOutCirc(t, b, c, d float64) float64 {
	t /= d / 2
	if t < 1 {
		return -c/2*(math.Sqrt(1-t*t)-1) + b
	}
	t -= 2
	return c/2*(math.Sqrt(1-t*t)+1) + b
}

//todo: EaseInElastic
//todo: EaseOutElastic
//todo: EaseInOutElastic
//todo: EaseInBack
//todo: EaseOutBack
//todo: EaseInOutBack
//todo: EaseInBounce
//todo: EaseOutBounce
//todo :EaseInOutBounce
