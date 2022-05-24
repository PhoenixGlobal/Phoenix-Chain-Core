#pragma once
/**
	@file
	@brief 192/256-bit additive homomorphic encryption by lifted-ElGamal
	@author MITSUNARI Shigeo(@herumi)
	@license modified new BSD license
	http://opensource.org/licenses/BSD-3-Clause
*/
#include <mcl/elgamal.hpp>
#include <mcl/ecparam.hpp>

namespace mcl {

#ifdef MCL_USE_AHE192
namespace ahe192 {

const mcl::EcParam& para = mcl::ecparam::NIST_P192;

typedef mcl::FpT<mcl::FpTag, 192> Fp;
typedef mcl::FpT<mcl::ZnTag, 192> Zn;
typedef mcl::EcT<Fp> Ec;
typedef mcl::ElgamalT<Ec, Zn> ElgamalEc;
typedef ElgamalEc::PrivateKey SecretKey;
typedef ElgamalEc::PublicKey PublicKey;
typedef ElgamalEc::CipherText CipherText;

static inline void initAhe()
{
	Fp::init(para.p);
	Zn::init(para.n);
	Ec::init(para.a, para.b);
	Ec::setIoMode(16);
	Zn::setIoMode(16);
}

static inline void initSecretKey(SecretKey& sec)
{
	const Ec P(Fp(para.gx), Fp(para.gy));
	sec.init(P, Zn::getBitSize());
}

} //mcl::ahe192
#endif

#ifdef MCL_USE_AHE256
namespace ahe256 {

const mcl::EcParam& para = mcl::ecparam::NIST_P256;

typedef mcl::FpT<mcl::FpTag, 256> Fp;
typedef mcl::FpT<mcl::ZnTag, 256> Zn;
typedef mcl::EcT<Fp> Ec;
typedef mcl::ElgamalT<Ec, Zn> ElgamalEc;
typedef ElgamalEc::PrivateKey SecretKey;
typedef ElgamalEc::PublicKey PublicKey;
typedef ElgamalEc::CipherText CipherText;

static inline void initAhe()
{
	Fp::init(para.p);
	Zn::init(para.n);
	Ec::init(para.a, para.b);
	Ec::setIoMode(16);
	Zn::setIoMode(16);
}

static inline void initSecretKey(SecretKey& sec)
{
	const Ec P(Fp(para.gx), Fp(para.gy));
	sec.init(P, Zn::getBitSize());
}

} //mcl::ahe256
#endif

} //  mcl
